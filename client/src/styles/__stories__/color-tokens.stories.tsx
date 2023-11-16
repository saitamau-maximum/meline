import type { Meta } from "@storybook/react";
import { vars } from "../theme.css";
import { Theme } from "./Theme";

const meta = {
    title: "Styles/Color Tokens",
} satisfies Meta;

export default meta;

export const ColorTokensLight = () => (
    <Theme.Light>
        <h1>Color Tokens</h1>
        <ul>
            {Object.entries(vars.color.gray).map(([key, value]) => (
                <>
                    <li key={key}>
                        <code>{key}</code>: <code>{value}</code>
                        <div
                            key={key}
                            style={{
                                backgroundColor: value,
                                height: "5rem",
                                width: "5rem",
                                borderRadius: "50%",
                            }}
                        />
                    </li>
                </>
            ))}
        </ul>
    </Theme.Light>
);

export const ColorTokensDark = () => (
    <Theme.Dark>
        <h1>Color Tokens</h1>
        <ul>
            {Object.entries(vars.color.gray).map(([key, value]) => (
                <>
                    <li key={key}>
                        <code>{key}</code>: <code>{value}</code>
                        <div
                            key={key}
                            style={{
                                backgroundColor: value,
                                height: "5rem",
                                width: "5rem",
                                borderRadius: "50%",
                            }}
                        />
                    </li>
                </>
            ))}
        </ul>
    </Theme.Dark>
);
