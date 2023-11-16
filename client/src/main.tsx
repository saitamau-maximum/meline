import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import { LoadingOverlayProvider } from "./providers/loading-overlay.tsx";
import { AuthProvider } from "./providers/auth.tsx";

import "@/styles/global.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <AuthProvider>
      <LoadingOverlayProvider>
        <App />
      </LoadingOverlayProvider>
    </AuthProvider>
  </React.StrictMode>
);
