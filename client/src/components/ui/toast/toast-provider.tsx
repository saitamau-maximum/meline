import { Toast } from "@/components/ui/toast";
import React, { createContext, useCallback, useState } from "react";
import { createPortal } from "react-dom";

interface ToastItem {
  id: string;
  children: React.ReactNode;
  to?: string;
}

interface ToastOptions {
  timeout: number;
  to: string;
}

interface ToastContextProps {
  pushToast: (
    children: React.ReactNode,
    options?: Partial<ToastOptions>
  ) => void;
}

export const ToastContext = createContext<ToastContextProps>({
  pushToast: () => {
    throw new Error("ToastContext is not implemented");
  },
});

interface LoadingProviderProps {
  children: React.ReactNode;
}

export const ToastProvider = ({ children }: LoadingProviderProps) => {
  const [items, setItems] = useState<ToastItem[]>([]);

  const pushToast = useCallback(
    (children: React.ReactNode, options: Partial<ToastOptions> = {}) => {
      const { timeout = 5000, to } = options;
      const id = crypto.randomUUID();
      setItems((prev) => [...prev, { id, children, to }]);
      setTimeout(() => {
        setItems((prev) => prev.filter((item) => item.id !== id));
      }, timeout);
    },
    []
  );

  return (
    <>
      <ToastContext.Provider
        value={{
          pushToast,
        }}
      >
        {children}
      </ToastContext.Provider>
      {createPortal(
        <Toast.Stack>
          {items.map((item) => (
            <Toast.Item key={item.id} to={item.to}>
              {item.children}
            </Toast.Item>
          ))}
        </Toast.Stack>,
        document.body
      )}
    </>
  );
};
