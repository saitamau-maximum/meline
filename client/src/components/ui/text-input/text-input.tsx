import { clsx } from "@/libs/clsx";
import { ComponentPropsWithRef, forwardRef } from "react";
import { styles } from "./text-input.css";

type Props = ComponentPropsWithRef<"input"> & {
  id: string;
  required?: boolean;
  label?: string;
  error?: string;
  description?: string;
};

export const TextInput = forwardRef<HTMLInputElement, Props>(
  ({ id, required, label, error, description, ...props }, ref) => (
    <div className={styles.textInputContainer}>
      {label && (
        <label className={styles.textInputLabel} htmlFor={id}>
          {label}
          {required && <span className={styles.textInputLabelRequired}>*</span>}
        </label>
      )}
      <input
        {...props}
        ref={ref}
        className={clsx(styles.textInput, error && styles.textInputError)}
        id={id}
      />
      {description && (
        <p className={styles.textInputDescription}>{description}</p>
      )}
      {error && <p className={styles.textInputErrorText}>{error}</p>}
    </div>
  )
);
