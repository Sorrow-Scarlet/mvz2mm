import { Moon, Sun, SunMoon } from "@lucide/svelte";
import type { Component } from "svelte";

export const darkOptions = [
  { id: "深色", icon: Moon },
  { id: "浅色", icon: Sun },
  { id: "自动", icon: SunMoon },
] as const satisfies ReadonlyArray<{ id: string; icon: Component }>;

export type DarkState = (typeof darkOptions)[number]["id"];

const order = ["深色", "浅色", "自动"];

export const currentDark = $state({ state: "自动" as DarkState });

export const changeDark = () => {
  const i = order.indexOf(currentDark.state);
  setDark(order[(i + 1) % order.length] as DarkState);
};

export const setDark = (next: DarkState) => {
  currentDark.state = next;
  localStorage.setItem("dark-mode", next);
};

const mq = window.matchMedia("(prefers-color-scheme: dark)");
export const systemDark = $state({ value: mq.matches });
mq.addEventListener("change", (e) => (systemDark.value = e.matches));

export function isDark(): boolean {
  return currentDark.state === "自动"
    ? systemDark.value
    : currentDark.state === "深色";
}
