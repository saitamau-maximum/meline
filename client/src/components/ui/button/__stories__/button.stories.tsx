import type { Meta } from "@storybook/react";
import { action } from "@storybook/addon-actions";
import { Button } from "..";

const meta = {
  title: "Components/UI/Button",
} satisfies Meta;

export default meta;

export const Primary = () => (
  <Button variant="primary" onClick={action("onClick")}>
    作成
  </Button>
);

export const Secondary = () => (
  <Button variant="secondary" onClick={action("onClick")}>
    一時保存
  </Button>
);

export const Text = () => (
  <Button variant="text" onClick={action("onClick")}>
    キャンセル
  </Button>
);

export const Overview = () => (
  <div style={{ display: "flex", gap: "1rem" }}>
    <Primary />
    <Secondary />
    <Text />
  </div>
);
