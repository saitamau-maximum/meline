import { LoadingOverlayContext } from "@/providers/loading-overlay";
import { useContext, useEffect } from "react";
import { useAuthUser } from "./auth-user";
import { useNotifications } from "@/features/notification/use-notifications";

export const useSetup = () => {
  const { setIsLoading } = useContext(LoadingOverlayContext);
  const { isLoading: isAuthUserLoading } = useAuthUser();
  useNotifications();

  useEffect(() => {
    setIsLoading(isAuthUserLoading);
  }, [isAuthUserLoading, setIsLoading]);
};
