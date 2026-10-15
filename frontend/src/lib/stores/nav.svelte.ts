import { LayoutGrid, Package, Globe, Settings } from "@lucide/svelte";
import type { Component } from "svelte";

export const navItems = [
  { title: "总览", url: "/dashboard", icon: LayoutGrid },
  { title: "Mod 管理", url: "/modmanage", icon: Package },
  { title: "网络源", url: "/netsource", icon: Globe },
  { title: "设置", url: "/settings", icon: Settings },
] as const satisfies ReadonlyArray<{
  title: string;
  url: string;
  icon: Component;
}>;
