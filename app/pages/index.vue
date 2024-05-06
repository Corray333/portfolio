<script setup>
import { onMounted } from 'vue'
import $ from 'jquery'


function fadeOut(el) {
    el.style.opacity = 1;
    (function fade() {
        if ((el.style.opacity -= .1) < 0) {
            el.style.display = "none";
        } else {
            requestAnimationFrame(fade);
        }
    })();
};

// ** FADE IN FUNCTION **
function fadeIn(el, display) {
    el.style.opacity = 0;
    el.style.display = display || "block";
    (function fade() {
        var val = parseFloat(el.style.opacity);
        if (!((val += .1) > 1)) {
            el.style.opacity = val;
            requestAnimationFrame(fade);
        }
    })();
};

const aboutMe = ref(new Map([
    ['Biography', ` My name is Mark.\nI was born in small city at the North of Russia. Now i live in Krasnodar and learn CS in Kuban Statu University.`],
    ['Education', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.'],
    ['Interests', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.'],
    ['Experience', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.'],
]))



/* SLIDE UP */
let slideUp = (target, pick, duration = 500) => {

    target.style.transitionProperty = 'height, margin, padding';
    target.style.transitionDuration = duration + 'ms';
    target.style.boxSizing = 'border-box';
    target.style.height = target.offsetHeight + 'px';
    target.offsetHeight;
    target.style.overflow = 'hidden';
    target.style.height = 0;
    target.style.paddingTop = 0;
    target.style.paddingBottom = 0;
    target.style.marginTop = 0;
    target.style.marginBottom = 0;
    window.setTimeout(() => {
        target.style.display = 'none';
        target.style.removeProperty('height');
        target.style.removeProperty('padding-top');
        target.style.removeProperty('padding-bottom');
        target.style.removeProperty('margin-top');
        target.style.removeProperty('margin-bottom');
        target.style.removeProperty('overflow');
        target.style.removeProperty('transition-duration');
        target.style.removeProperty('transition-property');

    }, duration);
}

/* SLIDE DOWN */
let slideDown = (target, duration = 500) => {

    target.style.removeProperty('display');
    let display = window.getComputedStyle(target).display;
    if (display === 'none') display = 'grid';
    target.style.display = display;
    let height = target.offsetHeight;
    target.style.overflow = 'hidden';
    target.style.height = 0;
    target.style.paddingTop = 0;
    target.style.paddingBottom = 0;
    target.style.marginTop = 0;
    target.style.marginBottom = 0;
    target.offsetHeight;
    target.style.boxSizing = 'border-box';
    target.style.transitionProperty = "height, margin, padding";
    target.style.transitionDuration = duration + 'ms';
    target.style.height = height + 'px';
    target.style.removeProperty('padding-top');
    target.style.removeProperty('padding-bottom');
    target.style.removeProperty('margin-top');
    target.style.removeProperty('margin-bottom');
    window.setTimeout(() => {
        target.style.removeProperty('height');
        target.style.removeProperty('overflow');
        target.style.removeProperty('transition-duration');
        target.style.removeProperty('transition-property');
    }, duration);
}

var togglePortfolio = (pick) => {
    if (pick === pickedProject.value) return
    const duration = 500
    const target = document.querySelector('.portfolio_card')

    target.style.transitionProperty = 'height, margin, padding';
    target.style.transitionDuration = duration + 'ms';
    target.style.boxSizing = 'border-box';
    target.style.height = target.offsetHeight + 'px';
    target.offsetHeight;
    target.style.overflow = 'hidden';
    target.style.height = 0;
    target.style.paddingTop = 0;
    target.style.paddingBottom = 0;
    target.style.marginTop = 0;
    target.style.marginBottom = 0;
    window.setTimeout(() => {
        target.style.display = 'none';
        target.style.removeProperty('height');
        target.style.removeProperty('padding-top');
        target.style.removeProperty('padding-bottom');
        target.style.removeProperty('margin-top');
        target.style.removeProperty('margin-bottom');
        target.style.removeProperty('overflow');
        target.style.removeProperty('transition-duration');
        target.style.removeProperty('transition-property');
        pickedProject.value = pick
        window.setTimeout(() => {
            slideDown(target)
        }, 50)
    }, duration);
}

const getAbout = (name) => {
    let aboutText = document.querySelector(`.about_text.${name}`)
    aboutMe.value.forEach((v, k) => {
        if (k !== name) {
            let aboutText = document.querySelector(`.about_text.${k}`)
            if (window.getComputedStyle(aboutText).display === 'block') {
                slideUp(aboutText)
            }
        }
    })
    if (window.getComputedStyle(aboutText).display === 'none') {
        slideDown(aboutText)
    }
}

const portfolioProjects = ref([
    {
        title: 'Stories service',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.',
        img: 'stories_service.png',
        stack: ['nonicons:vue-16', 'akar-icons:postgresql-fill', 'cib:redis', 'fa6-brands:golang', 'mdi:docker']
    },
    {
        title: 'Personal portfolio',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.',
        img: 'personal_portfolio.png',
        stack: ['akar-icons:postgresql-fill', 'fa6-brands:golang', 'mdi:docker', 'simple-icons:nuxtdotjs']
    },
    {
        title: 'Coding monopoly game',
        description: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi voluptatem, possimus voluptatibus officiis velit ipsa eum. Odio vel nam ipsa placeat nemo, ducimus tempore necessitatibus tempora, veniam assumenda laboriosam perferendis.',
        img: 'coding_monopoly_game.png',
        stack: ['nonicons:vue-16', 'fa6-brands:golang']
    }
])


const card = ref(null)

onMounted(() => {
    card.value = document.querySelector('.portfolio_card')
})

const rotate3D = (e) => {
    if (e.target.closest('.portfolio_card')) {
        card.value.style.transform = `scale(1.1)`
        return
    }
    var rect = card.value.getBoundingClientRect();
    let xAngle = (rect.height / 2 - (e.clientY - rect.top)) / 35
    let yAngle = (rect.width / 2 - (e.clientX - rect.left)) / 35
    card.value.style.transform = `rotateX(${xAngle}deg) rotateY(${yAngle}deg)`
}


const pickedProject = ref(portfolioProjects.value[0])

</script>
<template>
    <section class="home">
        <div
            class="route z-50 h-[33vh] flex flex-col items-center justify-between fixed px-5 my-auto top-[calc(50vh-33vh/2)]">
            <span class="line border-2 h-full border-[#1b1b1b] absolute py-1"></span>
            <a href="#header_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">Header</i>
            </a>
            <a href="#about_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">About
                    me</i>
            </a>
            <a href="#stack_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">My
                    stack</i>
            </a>
            <a href="#portfolio_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">Portfolio</i>
            </a>
            <a href="#blog_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">Blog</i>
            </a>
            <a href="#contact_container" class="point pointer group duration-200 hover:scale-[2]">
                <i
                    class="route_tip scale-0 opacity-50 group-hover:scale-100 group-hover:opacity-100 duration-200 origin-left">Contacts</i>
            </a>
        </div>
        <header id="header_container" class=" h-screen w-full flex flex-col items-center justify-center">
            <div class="title w-full flex flex-col items-center h-full justify-center">
                <svg width="60%" height="auto" viewBox="0 0 600 75">
                    <text style="font-family:'Courier Prime', monospace;" x="0" y="80%" font-size="75" fill="black">Hi,
                        It's Mark</text>
                </svg>
                <svg width="60%" height="auto" viewBox="0 0 600 75">
                    <text style="font-family:'Courier Prime', monospace;" x="0" y="50%" font-size="45" fill="black">and
                        I am WEB developer</text>
                </svg>
            </div>
            <img src="../public/images/home/header-hero.png" alt="" class=" z-20">
        </header>
        <section id="about_container" class="about w-screen h-screen p-20">
            <div class="dot one"></div>
            <div class="empty"></div>
            <div class="dot two"></div>
            <div class="empty"></div>
            <div class="dot three"></div>

            <div class="empty"></div>
            <h1 class="text-[4rem] font-bold">About me</h1>
            <div class="empty"></div>
            <div class="empty"></div>
            <div class="empty"></div>

            <div class="dot two"></div>
            <div class="empty"></div>
            <div class="dot three"></div>
            <div class="empty"></div>
            <div class="dot four"></div>

            <div class="empty"></div>
            <div class="empty"></div>
            <div class="empty"></div>
            <div class="empty"></div>
            <div class="empty"></div>


            <div class="dot three"></div>
            <div class="empty"></div>
            <div class="dot four"></div>
            <div class="empty"></div>
            <div class="dot five"></div>

            <div class="empty"></div>
            <ul class="about_list flex flex-col justify-between">
                <li v-for="(v, k) in aboutMe" :key="k" class="pointer btn"><a class=" pointer underline text-xl"
                        @click="getAbout(v[0])">{{ v[0]
                        }}</a></li>
            </ul>
            <div class="empty"></div>
            <article>
                <p style="display:none;" class=" about_text max-w-[36rem]" v-for="(v, k) in aboutMe" :key="k"
                    :class="v[0]">{{ v[1] }}</p>
            </article>
            <div class="empty"></div>

            <div class="dot four"></div>
            <div class="empty"></div>
            <div class="dot five"></div>
            <div class="empty"></div>
            <div class="dot six"></div>
        </section>

        <section id="stack_container" class="stack overflow-hidden py-20 w-full h-screen">
            <h1 class=" text-right px-20 text-[4rem]">My stack</h1>
            <div class="stack_wrapper w-full h-full flex justify-center items-center relative">
                <div class="stack_box">
                    <div
                        class=" pointer duration-300 hover:scale-110 vue flex justify-center items-center text-[8rem] backdrop-blur-md z-40">
                        <Icon name="nonicons:vue-16" />
                    </div>
                    <div
                        class=" pointer duration-300 hover:scale-110 postgres flex justify-center items-center text-[4rem] backdrop-blur-md z-40">
                        <Icon name="akar-icons:postgresql-fill" />
                    </div>
                    <div
                        class=" pointer duration-300 hover:scale-110 redis flex justify-center items-center text-[4rem] backdrop-blur-md z-40">
                        <Icon name="cib:redis" />
                    </div>
                    <div
                        class=" pointer duration-300 hover:scale-110 go flex justify-center items-center text-[8rem] backdrop-blur-md z-50">
                        <Icon name="fa6-brands:golang" />
                    </div>
                    <div
                        class=" pointer duration-300 hover:scale-110 docker flex justify-center items-center text-[4rem] backdrop-blur-md z-40">
                        <Icon name="mdi:docker" />
                    </div>
                    <div
                        class=" pointer duration-300 hover:scale-110 github flex justify-center items-center text-[8rem] backdrop-blur-md z-40">
                        <Icon name="logos:github" />
                    </div>
                    <div
                        class=" pointer duration-300  hover:scale-110 nuxt flex justify-center items-center text-[4rem] backdrop-blur-md z-40">
                        <Icon name="simple-icons:nuxtdotjs" />
                    </div>
                </div>
                <div class="stack_bg overflow-hidden absolute top-0 w-full h-full left-0 flex flex-col justify-between">
                    <h1 class="marquee1 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5">golang postgres
                        vue nuxt html css js mysql chi websockets redis docker nginx</h1>
                    <h1 class="marquee2 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5">js mysql chi
                        websockets redis docker nginx golang postgres vue nuxt html css</h1>
                    <h1 class="marquee3 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5">vue redis docker
                        nginx golang postgres nuxt html css js mysql chi websockets js</h1>
                    <h1 class="marquee4 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5">golang postgres
                        nuxt html vue redis js mysql chi websockets js docker nginx css</h1>
                    <h1 class="marquee5 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5">mysql chi
                        websockets js golang postgres nuxt html vue redis docker nginx css js</h1>
                    <h1 class="marquee3 whitespace-nowrap text-[10rem] uppercase font-bold opacity-5"> postgres nuxt
                        html css vue redis docker nginx golang chi websockets js js mysql</h1>
                </div>
            </div>
        </section>

        <section id="portfolio_container" @mousemove="rotate3D"
            class="portfolio w-full h-screen relative overflow-hidden flex p-20">
            <div class="portfolio_card_container overflow-hidden w-full h-full flex justify-center items-center">
                <PortfolioCard :project="pickedProject" />
            </div>
            <ul class="flex flex-col items-end">
                <li v-for="(project, i) of portfolioProjects" :key="i" @click="togglePortfolio(project)"
                    class="btn w-fit text-right pointer lowercase underline whitespace-nowrap"><a>{{ project.title
                        }}</a></li>
            </ul>
            <p class=" z-0 text-left text-[25vw] absolute font-bold uppercase opacity-5 left-0 bottom-0 line-h-fit">
                Port<br />folio</p>
        </section>
        <section id="blog_container" class="blog flex flex-col w-screen h-screen p-20">
            <h1 class="text-[4rem]">Blog</h1>
            <section class="featured_posts w-full h-full flex justify-center items-center">
                <p class=" text-4xl">Coming soon</p>
            </section>
        </section>

        <div class="bottom w-full h-screen grid grid-rows-4 grid-cols-1">
            <section id="contact_container" class=" row-span-3 w-full h-full p-20 ">
                <div class="content flex flex-col    items-center justify-center h-full">
                    <h1 class="text-[4rem] text-center">Contact me</h1>
                    <div class="container flex items-center justify-center gap-3">
                        <div class="contact_form flex flex-col gap-4">
                            <span class="flex flex-col  ">
                                <label for="">How to contact you:</label>
                                <p class="inpt z-50 pointer">
                                    <a>
                                        <input type="email" placeholder="email"
                                            class="   z-50 border-b-2 border-dark max-w-[20rem]">
                                    </a>

                                </p>
                            </span>
                            <span class="flex flex-col  ">
                                <label for="">Message:</label>
                                <p class="inpt z-50 pointer">
                                    <a>
                                        <input type="text" placeholder="your message"
                                            class="   z-50 border-b-2 border-dark max-w-[20rem]">
                                    </a>

                                </p>
                            </span>
                        </div>
                        <span class="h-32 w-[2px] bg-dark"></span>
                        <div class="contact_networks flex flex-col items-center">
                            <h1>Or use my social networks:</h1>
                            <div class="flex text-[3rem] z-50">
                                <Icon class="pointer duration-300 hover:bg-dark hover:text-white p-2 rounded-full"
                                    name="mdi:telegram" />
                                <Icon class="pointer duration-300 hover:bg-dark hover:text-white p-2 rounded-full"
                                    name="mdi:github" />
                            </div>
                        </div>
                    </div>

                </div>
            </section>
            <footer class="bg-dark w-full h-full flex justify-center items-center text-white">
                Designed and developed by Mark Anikin 2024
            </footer>
        </div>
    </section>
</template>


<style scoped>
.title {
    font-family: 'Courier Prime', monospace;
    line-height: 0.8;
}


/* TODO: change */
.inpt>a {
    position: relative;
    width: fit-content;
    background-color: dark;
}

.inpt>a::after {
    content: '';
    background-color: white;
    width: 100%;
    height: 100%;
    transform: scaleX(0);
    transform-origin: left;
    position: absolute;
    top: 0;
    left: 0;
    mix-blend-mode: difference;
    transition: transform 0.2s;
}

.inpt:hover a::after {
    transform: scaleX(1);
}
.inpt:hover input{
    color:black !important;
}

.title>svg {
    height: fit-content;
}

.point {
    width: 1rem;
    height: 1rem;
    background-color: #1b1b1b;
    border-radius: 100%;
    display: flex;
    align-items: center;
    line-height: 1;
    z-index: 11;
}

.route_tip {
    margin-left: 2rem;
    white-space: nowrap;
    background-color: #fff;
}

.about {
    display: grid;
    grid-template-columns: 2rem 1fr 2rem 1fr 2rem;
    grid-template-rows: 2rem 1fr 2rem 1fr 2rem 1fr 2rem;
}

.about_list>li {
    width: fit-content;
}

.btn>a {
    position: relative;
    width: fit-content;
    background-color: white;
}


.about_text {
    white-space: pre-line;
}

.dot {
    width: 100%;
    height: 100%;
    background-color: #1b1b1b;
    animation: infinite 2s dotWave;
}

.dot.one {
    animation-delay: 0s;
}

.dot.two {
    animation-delay: 0.2s;
}

.dot.three {
    animation-delay: 0.4s;
}

.dot.four {
    animation-delay: 0.6s;
}

.dot.five {
    animation-delay: 0.8s;
}

.dot.six {
    animation-delay: 1s;
}

@keyframes dotWave {
    0% {
        transform: scale(1);
    }

    50% {
        transform: scale(0.6);
    }

    100% {
        transform: scale(1);
    }
}

.stack_bg>h1 {
    font-family: 'Roboto Mono', monospace;
    line-height: 0.8;
}


.marquee1 {
    animation: marquee 38s linear infinite;
}

.marquee2 {
    animation: marquee 40s linear infinite;
}

.marquee3 {
    animation: marquee 35s linear infinite;
}

.marquee4 {
    animation: marquee 44s linear infinite;
}

.marquee5 {
    animation: marquee 36s linear infinite;
}

@keyframes marquee {
    0% {
        transform: translateX(0);
    }

    50% {
        transform: translateX(-100%);
    }

    100% {
        transform: translateX(0);
    }
}


.stack_box {
    display: grid;
    grid-template-columns: repeat(4, 9rem);
    grid-template-rows: repeat(3, 9rem);
    gap: 1rem
}

.stack_box>div {
    background-color: #F3F3F3;
    opacity: 0.8;
    box-shadow: 0 0 25px 0 rgba(0, 0, 0, 0.25);
    border-radius: 25px;
    width: 100%;
    height: 100%;
}

.stack_box>div>* {
    z-index: 1;
}

.stack_box>div>* {
    z-index: 20;
}

.stack_box>div::before {
    content: "";
    position: absolute;
    top: -2px;
    left: -2px;
    right: -2px;
    bottom: -2px;
    background: linear-gradient(135deg, #ffffff, rgb(221, 221, 221));
    border-radius: 30px;
    /* 25px + border-width */
    z-index: 9;
}

.stack_box>div::after {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: #F3F3F3;
    backdrop-filter: blur(10px);
    box-shadow: inset 0 0 10px 2px rgb(27, 27, 27, 0.05);
    border-radius: 25px;
    z-index: 10;
}

.vue {
    grid-column: 1/ span 2;
    grid-row: 1 / span 2;
    width: 100%;
    height: 100%;
    z-index: 60;

}

.go {
    grid-column: 3/ span 2;
    grid-row: 1;
    width: 100%;
    height: 100%;
}

.github {
    grid-column: 2/span 2;
}
</style>