import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense } from "react";

function App() {

  return (
    <>
      <LoadingOverlay />
      <BrowserRouter>
        <Suspense>
          <AppRoutes />
        </Suspense>
      </BrowserRouter>
    </>
  );
}

export default App;
