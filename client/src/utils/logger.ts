import { CSSProperties } from "react";

const loggerStyle: CSSProperties = {
  fontWeight: "bold",
  color: "white",
  padding: "2px 4px",
  borderRadius: "4px",
};

const encodeStyle = (style: CSSProperties) => {
  return Object.entries(style)
    .map(
      ([key, value]) =>
        `${key.replace(/([A-Z])/g, "-$1").toLowerCase()}:${value}`
    )
    .join(";");
};

const loggerVariant = {
  INFO: {
    background: "green",
  },
  WARN: {
    background: "orange",
  },
  ERROR: {
    background: "red",
  },
} as const;

const baseLogger = (message: string, variant: keyof typeof loggerVariant) => {
  console.log(
    `%c${variant}`,
    `${encodeStyle(loggerStyle)};${encodeStyle(loggerVariant[variant])}`,
    message
  );
};

export const logger = {
  info: (message: string) => baseLogger(message, "INFO"),
  warn: (message: string) => baseLogger(message, "WARN"),
  error: (message: string) => baseLogger(message, "ERROR"),
  raw: console.log,
};
