import { LayoutGrid, Package, Globe, Settings } from "@lucide/svelte";
import type { Component } from "svelte";

export const navItems = [
  { id: "dashboard", title: "总览", icon: LayoutGrid },
  { id: "modmanage", title: "Mod 管理", icon: Package },
  { id: "netsource", title: "网络源", icon: Globe },
  { id: "settings", title: "设置", icon: Settings },
] as const satisfies ReadonlyArray<{
  id: string;
  title: string;
  icon: Component;
}>;

export type PageID = (typeof navItems)[number]["id"];

export const nav = $state({ page: "dashboard" as PageID });
