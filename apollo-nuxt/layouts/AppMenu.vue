<script setup>
import { ref } from 'vue';

import AppSubMenu from './AppSubMenu.vue';

const model = ref([
    {
        label: 'Pages',
        icon: 'pi pi-fw pi-briefcase',
        items: [
            {
                label: 'Landing',
                icon: 'pi pi-fw pi-globe',
                to: '/landing/'
            },
            {
                label: 'Crud',
                icon: 'pi pi-fw pi-pencil',
                to: '/pages/crud'
            },
            {
                label: 'Invoice',
                icon: 'pi pi-fw pi-dollar',
                to: '/pages/invoice'
            },
            {
                label: 'About Us',
                icon: 'pi pi-fw pi-user',
                to: '/pages/aboutus'
            },
            {
                label: 'Help',
                icon: 'pi pi-fw pi-question-circle',
                to: '/pages/help'
            },
            {
                label: 'FAQ',
                icon: 'pi pi-fw pi-question',
                to: '/pages/faq'
            },
            {
                label: 'Contact Us',
                icon: 'pi pi-fw pi-phone',
                to: '/pages/contact'
            }
        ]
    }
]);
</script>

<template>
    <AppSubMenu :model="model" />
</template>

<style lang="scss" scoped></style>
