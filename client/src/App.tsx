import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { LoadingOverlayContext } from "./providers/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense, useContext, useEffect } from "react";

function App() {
  // const { setIsLoading } = useContext(LoadingOverlayContext);

  // useEffect(() => {
  //   setIsLoading(true);

  //   setTimeout(() => {
  //     setIsLoading(false);
  //   }, 3000);
  // }, [setIsLoading]);

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
