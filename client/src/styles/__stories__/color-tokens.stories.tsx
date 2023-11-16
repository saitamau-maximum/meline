import type { Meta } from "@storybook/react";
import { vars } from "..";
import { Theme } from "./Theme";

const meta = {
  title: "Styles/Color Tokens",
} satisfies Meta;

export default meta;

const ColorTokens = () => (
  <>
    <h1>Color Tokens</h1>
    <div
      style={{
        display: "flex",
        gap: "32px",
        padding: "32px",
      }}
    >
      {Object.entries(vars.color.gray).map(([key, value]) => (
        <div
          key={key}
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            gap: "12px",
          }}
        >
          <div
            key={key}
            style={{
              backgroundColor: value,
              height: "48px",
              width: "48px",
              borderRadius: "50%",
              border: `1px solid ${vars.color.gray[6]}`,
            }}
          />
          <code
            style={{
              fontSize: "0.75rem",
              color: vars.color.gray[12],
            }}
          >
            Gray {key}
          </code>
        </div>
      ))}
    </div>
  </>
);

export const Light = () => (
  <Theme.Light>
    <ColorTokens />
  </Theme.Light>
);

export const Dark = () => (
  <Theme.Dark>
    <ColorTokens />
  </Theme.Dark>
);
