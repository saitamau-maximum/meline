import * as D from "@radix-ui/react-dialog";
import { ComponentProps } from "react";
import { styles } from "./dialog.css";
import { clsx } from "@/libs/clsx";
import { Button } from "../button";
import { X } from "react-feather";

export const Dialog = {
  Root: D.Root,
  Trigger: D.Trigger,
  Portal: D.Portal,
  Overlay: ({ className, ...props }: ComponentProps<typeof D.Overlay>) => (
    <D.Overlay className={clsx(styles.dialogOverlay, className)} {...props} />
  ),
  Content: ({ className, ...props }: ComponentProps<typeof D.Content>) => (
    <D.Content className={clsx(styles.dialogContent, className)} {...props} />
  ),
  Title: ({ className, ...props }: ComponentProps<typeof D.Title>) => (
    <D.Title className={clsx(styles.dialogTitle, className)} {...props} />
  ),
  Description: ({
    className,
    ...props
  }: ComponentProps<typeof D.Description>) => (
    <D.Description
      className={clsx(styles.dialogDescription, className)}
      {...props}
    />
  ),
  RightTopClose: (props: ComponentProps<typeof D.Close>) => (
    <Dialog.Close asChild {...props}>
      <Button
        variant="text"
        size="sm"
        className={styles.dialogRightTopCloseButton}
      >
        <X />
      </Button>
    </Dialog.Close>
  ),
  Footer: ({ className, ...props }: ComponentProps<"div">) => (
    <div className={clsx(styles.dialogFooter, className)} {...props} />
  ),
  Close: D.Close,
};
