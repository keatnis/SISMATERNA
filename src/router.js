import Vue from 'vue'
import Router from 'vue-router'
import Mujer from '@/components/Mujer'
import RegistrarEmbarazada from '@/components/RegistrarEmbarazada'
import Ajustes from '@/components/Ajustes'

import Informes from '@/components/Informes'

import Escritorio from '@/components/Escritorio'

import Login from '@/components/Login'
import Cobrar from '@/components/Cobrar'


Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Mujer',
            component: Mujer
        },
        {
            path: '/agregar-embarazada',
            name: 'RegistrarEmbarazada',
            component: RegistrarEmbarazada,
        },
        {
            path: '/ajustes',
            name: 'Ajustes',
            component: Ajustes,
        },
        {
            path: '/informes',
            name: 'Informes',
            component: Informes,
        },
        {
            path: '/escritorio',
            name: 'Escritorio',
            component: Escritorio,
        },
        
        {
            path: '/Login',
            name: 'Login',
            component: Login
        },
        {
            path: '/Cobrar',
            name: 'Cobrar',
            component: Cobrar
        }

    ]
});