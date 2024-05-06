<script setup>




</script>

<template>
    <section class=" p-20 flex flex-col text-2xl justify-end text-white bg-transparent w-screen h-screen z-50">
        <ul>
            <li class="btn pointer">
                <nuxt-link>home</nuxt-link>
            </li>
            <li class="btn pointer">
                <nuxt-link>blog</nuxt-link>
            </li>
            <li class="btn pointer">
                <nuxt-link>contacts</nuxt-link>
            </li>
        </ul>
    </section>
</template>


<style scoped>
@layer navbar {
    ul {
        @apply underline flex flex-col gap-5 w-fit;
    }
    ul>li{
        @apply w-fit;
    }

    

}
</style>