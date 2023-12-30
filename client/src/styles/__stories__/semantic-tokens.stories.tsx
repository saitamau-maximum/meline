import type { Meta } from "@storybook/react";
import { vars } from "..";

const meta = {
  title: "Styles/Semantic Tokens",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <>
    <h1>Semantic Tokens</h1>
    <h2>Background</h2>
    <div
      style={{
        display: "flex",
        gap: "32px",
        padding: "32px",
      }}
    >
      {Object.entries(vars.semantic.background).map(([key, value]) => (
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
    <h2>Text</h2>
    <div
      style={{
        display: "flex",
        gap: "32px",
        padding: "32px",
      }}
    >
      {Object.entries(vars.semantic.text).map(([key, value]) => (
        <div
          key={key}
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            gap: "12px",
          }}
        >
          <span
            key={key}
            style={{
              color: value,
            }}
          >
            This is {key}
          </span>
        </div>
      ))}
    </div>
  </>
);
