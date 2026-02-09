/**
 * Grafana Dashboard Embed Card — KubeStellar Console integration
 * Add to web/src/components/cards/ in your console fork.
 *
 * Props: dashboardUid, from, to
 */

import React from 'react';

interface GrafanaDashboardCardProps {
  dashboardUid: string;
  from?: string;
  to?: string;
}

export const GrafanaDashboardCard: React.FC<GrafanaDashboardCardProps> = ({
  dashboardUid,
  from = 'now-1h',
  to = 'now',
}) => {
  const embedUrl = `/api/integrations/grafana/embed?uid=${dashboardUid}&from=${encodeURIComponent(from)}&to=${encodeURIComponent(to)}`;

  return (
    <div className="rounded-lg border overflow-hidden">
      <div className="p-2 border-b text-sm font-medium">Grafana: {dashboardUid}</div>
      <iframe
        src={embedUrl}
        title={`Grafana dashboard ${dashboardUid}`}
        className="w-full h-[400px] border-0"
        sandbox="allow-scripts allow-same-origin"
      />
    </div>
  );
};
