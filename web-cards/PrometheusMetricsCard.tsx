/**
 * Prometheus Metrics Card — KubeStellar Console integration
 * Add to web/src/components/cards/ in your console fork.
 *
 * Props: cluster, metric (e.g. container_cpu_usage_seconds_total), timeRange
 */

import React, { useEffect, useState } from 'react';

interface PrometheusMetricsCardProps {
  cluster?: string;
  metric?: string;
  timeRange?: string;
}

export const PrometheusMetricsCard: React.FC<PrometheusMetricsCardProps> = ({
  cluster = '',
  metric = 'container_cpu_usage_seconds_total',
  timeRange = '1h',
}) => {
  const [data, setData] = useState<{ timestamp: string; value: number }[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const params = new URLSearchParams({ cluster, metric, range: timeRange });
    fetch(`/api/integrations/prometheus/metrics?${params}`)
      .then((res) => res.json())
      .then((d) => {
        setData(d.data || []);
        setLoading(false);
      })
      .catch((e) => {
        setError(e.message);
        setLoading(false);
      });
  }, [cluster, metric, timeRange]);

  if (loading) return <div className="p-4">Loading metrics...</div>;
  if (error) return <div className="p-4 text-red-500">Error: {error}</div>;

  return (
    <div className="rounded-lg border p-4">
      <h3 className="mb-2 font-semibold">Prometheus: {metric}</h3>
      <div className="min-h-[200px]">
        {data.length === 0 ? (
          <p className="text-muted-foreground">No data. Configure PROMETHEUS_URL and ensure the query returns results.</p>
        ) : (
          <pre className="text-xs overflow-auto">{JSON.stringify(data.slice(0, 5), null, 2)}...</pre>
        )}
      </div>
    </div>
  );
};
