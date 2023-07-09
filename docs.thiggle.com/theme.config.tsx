import { DocsThemeConfig } from "nextra-theme-docs";
import Wordmark from "./components/Wordmark";

const config: DocsThemeConfig = {
  logo: (
    <span style={{ display: "flex", gap: "15px" }}>
      <Wordmark /> docs
    </span>
  ),
  project: {
    link: "https://docs.thiggle.com",
  },
  docsRepositoryBase: "https://github.com/thiggle/api",
  feedback: {
    labels: "area/docs",
  },
  footer: {
    text: "All rights reserved.",
  },
  useNextSeoProps() {
    return {
      titleTemplate: "%s – Thiggle Docs",
    };
  },
};

export default config;
