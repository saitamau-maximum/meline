import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense, useCallback, useEffect } from "react";
import { useAuth } from "./providers/auth";
import { object, string, safeParse } from "valibot";

const UserMeResponseSchema = object({
  name: string(),
  image_url: string(),
});

function App() {
  const { setUser } = useAuth();
  const initializeUser = useCallback(async () => {
    const res = await fetch("/api/users/me");
    if (!res.ok) throw new Error("Unauthorized");
    const user = safeParse(UserMeResponseSchema, await res.json());
    if (!user.success) throw new Error("Invalid response");
    setUser({
      name: user.output.name,
      imageURL: user.output.image_url,
    });
  }, [setUser]);

  useEffect(() => {
    void initializeUser();
  }, [initializeUser]);

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
