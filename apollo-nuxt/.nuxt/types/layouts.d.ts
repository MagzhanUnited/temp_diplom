import { ComputedRef, MaybeRef } from 'vue'
export type LayoutKey = "app-breadcrumb" | "app-config" | "app-layout" | "app-menu" | "app-menu-item" | "app-profile-sidebar" | "app-sidebar" | "app-sub-menu" | "app-topbar" | "composables-layout" | "default"
declare module "../../node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    layout?: MaybeRef<LayoutKey | false> | ComputedRef<LayoutKey | false>
  }
}