export default defineNuxtRouteMiddleware((to, from) => {
    // const a =
    if (process.server) return;

    if (to.path !== '/auth/login' && !localStorage.getItem('token')) {
        return navigateTo('/auth/login');
    }
});
