import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";

import "@/styles/global.css";

const enableMocking = async () => {
  if (import.meta.env.MODE !== "msw") return;
  const { worker } = await import("./mocks/browser");
  return worker.start({
    onUnhandledRequest: "bypass",
  });
};

await enableMocking().then(() => {
  ReactDOM.createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
      <App />
    </React.StrictMode>
  );
});
