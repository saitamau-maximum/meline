import { clsx } from "@/libs/clsx";
import { ComponentPropsWithRef, forwardRef } from "react";
import { styles } from "./textarea.css";

type Props = ComponentPropsWithRef<"textarea"> & {
  id: string;
  required?: boolean;
  label?: string;
  error?: string;
  description?: string;
};

export const Textarea = forwardRef<HTMLTextAreaElement, Props>(
  ({ id, required, label, error, description, ...props }, ref) => (
    <div className={styles.textAreaContainer}>
      {label && (
        <label className={styles.textAreaLabel} htmlFor={id}>
          {label}
          {required && <span className={styles.textAreaLabelRequired}>*</span>}
        </label>
      )}
      <textarea
        {...props}
        ref={ref}
        className={clsx(styles.textArea, error && styles.textAreaError)}
        id={id}
      />
      {description && (
        <p className={styles.textAreaDescription}>{description}</p>
      )}
      {error && <p className={styles.textAreaErrorText}>{error}</p>}
    </div>
  )
);
