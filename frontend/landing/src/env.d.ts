/// <reference path="../.astro/types.d.ts" />

interface ImportMetaEnv {
  readonly PUBLIC_WP_HOST: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
