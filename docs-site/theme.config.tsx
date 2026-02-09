import React from 'react'
import { DocsThemeConfig } from 'nextra-theme-docs'

const config: DocsThemeConfig = {
  logo: <span><strong>KubeStellar Console Integrations</strong></span>,
  project: {
    link: 'https://github.com/kubestellar/kss-demo',
  },
  docsRepositoryBase: 'https://github.com/kubestellar/kss-demo/tree/main/docs-site',
  footer: {
    text: 'KubeStellar Console Integrations — CNCF LFX Mentorship 2026 Term 1',
  },
  useNextSeoProps() {
    return {
      titleTemplate: '%s – KubeStellar Console Integrations',
    }
  },
}

export default config
