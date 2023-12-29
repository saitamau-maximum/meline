import { vars as vs } from "@/styles/theme.css";

export const vars = {
  ...vs,
  transition: {
    normal: (...properties: string[]) => {
      return properties
        .map((property) => `${property} 0.3s ease-in-out`)
        .join(", ");
    },
  },
};

export const constants = {
  sizes: {
    channelLayoutSidePanelWidth: "192px",
  },
};
