<script setup>

const switchLanguage = ()=>{
    if (localStorage.getItem('lang') === 'ru') {
        localStorage.setItem('lang', 'eng')
    } else {
        localStorage.setItem('lang', 'ru')
    }
    location.reload()
}


</script>

<template>
    <section class=" p-20 flex flex-col text-2xl justify-end text-white bg-transparent w-screen h-screen z-50">
        <ul>
            <li class="btn pointer">
                <nuxt-link to="/" @click="$emit('route')">home</nuxt-link>
            </li>
            <li class="btn pointer">
                <nuxt-link to="/blog" @click="$emit('route')">blog</nuxt-link>
            </li>
            <li class="btn pointer">
                <nuxt-link @click="$emit('route')">contacts</nuxt-link>
            </li>
            <li>
                <div class="btn"><a class="bg-dark pointer" @click="switchLanguage">switch blog language</a></div>
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