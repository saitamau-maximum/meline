import { ComponentProps, forwardRef } from "react";
import { styles } from "./button.css";
import { clsx } from "@/libs/clsx";

interface ButtonProps extends ComponentProps<"button"> {
  variant?: "primary" | "secondary" | "text";
  size?: "sm" | "md";
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ variant = "primary", size = "md", className, ...props }, ref) => (
    <button
      ref={ref}
      className={clsx(
        styles.button,
        styles.buttonVariant[variant],
        styles.buttonSize[size],
        className
      )}
      {...props}
    />
  )
);
