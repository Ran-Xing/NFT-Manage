import {createRouter,createWebHashHistory} from "vue-router";
import NFT from "@/components/NFT";
import Home from "@/components/HomePage";

const router= createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            component: Home
        },
        {
            path: "/nft",
            component: NFT
        }
    ]
});

export default router;
