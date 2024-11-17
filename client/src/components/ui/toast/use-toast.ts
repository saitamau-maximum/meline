import { ToastContext } from "./toast-provider";
import { useContext } from "react";

export const useToast = () => {
  return useContext(ToastContext);
};
