/**
 * GitHub Workflows Card — KubeStellar Console integration
 * Add to web/src/components/cards/ in your console fork.
 *
 * Props: owner, repo
 */

import React, { useEffect, useState } from 'react';

interface WorkflowRun {
  id: number;
  name: string;
  status: string;
  conclusion: string;
  html_url: string;
  created_at: string;
}

interface GitHubWorkflowsCardProps {
  owner?: string;
  repo?: string;
}

export const GitHubWorkflowsCard: React.FC<GitHubWorkflowsCardProps> = ({
  owner = 'kubestellar',
  repo = 'console',
}) => {
  const [workflows, setWorkflows] = useState<WorkflowRun[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const params = new URLSearchParams({ owner, repo });
    fetch(`/api/integrations/github/workflows?${params}`)
      .then((res) => res.json())
      .then((d) => {
        setWorkflows(d.workflow_runs || []);
        setLoading(false);
      })
      .catch((e) => {
        setError(e.message);
        setLoading(false);
      });
  }, [owner, repo]);

  if (loading) return <div className="p-4">Loading workflows...</div>;
  if (error) return <div className="p-4 text-red-500">Error: {error}</div>;

  return (
    <div className="rounded-lg border p-4">
      <h3 className="mb-2 font-semibold">GitHub Workflows: {owner}/{repo}</h3>
      <ul className="space-y-2">
        {workflows.slice(0, 10).map((wf) => (
          <li key={wf.id} className="flex items-center gap-2 text-sm">
            <span
              className={`w-2 h-2 rounded-full ${
                wf.status === 'completed'
                  ? wf.conclusion === 'success'
                    ? 'bg-green-500'
                    : 'bg-red-500'
                  : 'bg-yellow-500'
              }`}
            />
            <a href={wf.html_url} target="_blank" rel="noopener noreferrer" className="hover:underline">
              {wf.name}
            </a>
            <span className="text-muted-foreground">
              {wf.status} {wf.conclusion && `(${wf.conclusion})`}
            </span>
          </li>
        ))}
      </ul>
    </div>
  );
};
