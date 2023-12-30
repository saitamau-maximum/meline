import type { Meta } from "@storybook/react";
import { action } from "@storybook/addon-actions";
import { Dialog } from "..";
import { Button } from "../../button";

const meta = {
  title: "Components/UI/Dialog",
} satisfies Meta;

export default meta;

export const Default = () => (
  <Dialog.Root>
    <Dialog.Trigger asChild>
      <Button onClick={action("onClick")}>ダイアログを開く</Button>
    </Dialog.Trigger>
    <Dialog.Overlay>
      <Dialog.Content>
        <Dialog.RightTopClose />
        <Dialog.Title>Dialog Title</Dialog.Title>
        <Dialog.Description>Dialog Description</Dialog.Description>
        <p style={{ textAlign: "right" }}>
          <Dialog.Close asChild>
            <Button variant="secondary">Cancel</Button>
          </Dialog.Close>
          <span style={{ margin: "0 0.5rem" }} />
          <Dialog.Close asChild>
            <Button variant="primary">OK</Button>
          </Dialog.Close>
        </p>
      </Dialog.Content>
    </Dialog.Overlay>
  </Dialog.Root>
);
