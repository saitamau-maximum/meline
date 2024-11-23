import { clsx } from "@/libs/clsx";
import { styles } from "./toast.css";
import { Link } from "react-router-dom";

interface ToastItemProps {
  children: React.ReactNode;
  to?: string;
}

const ToastItem = ({ children, to }: ToastItemProps) => {
  if (to) {
    return (
      <Link to={to} className={clsx(styles.toastItemContainer, styles.link)}>
        {children}
      </Link>
    );
  }

  return <div className={styles.toastItemContainer}>{children}</div>;
};

interface ToastStackProps {
  children: React.ReactNode;
}

const ToastStack = ({ children }: ToastStackProps) => {
  return <div className={styles.toastStackContainer}>{children}</div>;
};

export const Toast = {
  Item: ToastItem,
  Stack: ToastStack,
};
