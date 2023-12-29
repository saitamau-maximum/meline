import { CHANNEL_TOP_ROUTE } from "@/features/channel/routes";
import { TOP_ROUTE } from "@/features/onboarding/routes";
import { AuthContext } from "@/providers/auth";
import { useContext } from "react";
import { Navigate } from "react-router-dom";

interface AuthProtectionWrapperProps {
  children: React.ReactNode;
}

export const AuthRequired = ({ children }: AuthProtectionWrapperProps) => {
  const { state } = useContext(AuthContext);

  if (!state.isAuthenticated) {
    return <Navigate to={TOP_ROUTE} />;
  }

  return <>{children}</>;
};

export const UnAuthRequired = ({ children }: AuthProtectionWrapperProps) => {
  const { state } = useContext(AuthContext);

  if (state.isAuthenticated) {
    return <Navigate to={CHANNEL_TOP_ROUTE} />;
  }

  return <>{children}</>;
};
