export type PageID = "dashboard" | "modmanage" | "netsource" | "settings";

export const nav = $state({ page: "dashboard" as PageID });
