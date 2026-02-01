import React from 'react'
import { DocsThemeConfig } from 'nextra-theme-docs'

const config: DocsThemeConfig = {
  logo: <span><strong>KubeStellar Integrations</strong></span>,
  project: {
    link: 'https://github.com/kubestellar/ks-demo',
  },
  docsRepositoryBase: 'https://github.com/kubestellar/ks-demo/tree/main/docs-site',
  footer: {
    text: 'KubeStellar Integrations Demo — CNCF LFX Mentorship 2026 Term 1',
  },
  useNextSeoProps() {
    return {
      titleTemplate: '%s – KubeStellar Integrations',
    }
  },
}

export default config
