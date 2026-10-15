import { Moon, Sun, SunMoon } from "@lucide/svelte";
import type { Component } from "svelte";

export const darkOptions = [
  { id: "enabled", icon: Moon },
  { id: "disabled", icon: Sun },
  { id: "auto", icon: SunMoon },
] as const satisfies ReadonlyArray<{ id: string; icon: Component }>;

export type DarkState = (typeof darkOptions)[number]["id"];

export const currentDark = $state({ state: "auto" as DarkState });
