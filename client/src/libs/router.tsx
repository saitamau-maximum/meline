import { useAuthUser } from "@/hooks/auth-user";
import { Navigate } from "react-router-dom";

interface AuthProtectionWrapperProps {
  children: React.ReactNode;
}

export const AuthRequired = ({ children }: AuthProtectionWrapperProps) => {
  const { isAuthenticated } = useAuthUser();

  if (!isAuthenticated) {
    return <Navigate to="/" />;
  }

  return <>{children}</>;
};

export const UnAuthRequired = ({ children }: AuthProtectionWrapperProps) => {
  const { isAuthenticated } = useAuthUser();

  if (isAuthenticated) {
    return <Navigate to="/channel" />;
  }

  return <>{children}</>;
};
