import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { vanillaExtractPlugin } from "@vanilla-extract/vite-plugin";

// https://vitejs.dev/config/
export default defineConfig((env) => ({
  plugins: [react(), vanillaExtractPlugin()],
  resolve: {
    alias: {
      "@": "/src",
    },
  },
  server: {
    port: 3000,
    hmr: {
      clientPort: env.mode === "msw" ? 3000 : 80,
      protocol: "ws",
    },
  },
}));
