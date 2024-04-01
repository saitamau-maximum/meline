import { vars as vs } from "@/styles/theme.css";

export const vars = {
  ...vs,
  transition: {
    normal: (...properties: string[]) => {
      return properties
        .map((property) => `${property} 0.3s ease-in-out`)
        .join(", ");
    },
    fastInteraction: (...properties: string[]) => {
      return properties
        .map((property) => `${property} 0.15s cubic-bezier(0.16, 1, 0.3, 1)`)
        .join(", ");
    },
  },
};

export const constants = {
  sizes: {
    channelLayoutSidePanelWidth: "240px",
    dialogMaxWidth: "450px",
  },
};
