<template>
    <!-- <div class="columns"> -->
    <div class="columns">
        <div class="column is-half centered">
            <div class="has-text-centered">
                <article class="panel  centered">
                    <div class="panel-heading">Registrar Usuario</div>
                </article>
            </div>
            <form action="..." @submit="checkForm" novalidate="true">
                <b-field label="No. Progreso">
                    <b-input disabled="»disabled»" v-model="detalles.id" placeholder=""></b-input>
                </b-field>
                <b-field label="Clues">
                    <b-input v-model="detalles.clues" placeholder="Clues"></b-input>
                </b-field>
                <b-field label="Nombre de la unidad">
                    <b-input v-model="detalles.nombreunidad" placeholder="Nombre de la unidad"></b-input>
                </b-field>
                <b-field label="Usuario" name="usuario">
                    <b-input v-model="detalles.usuario" placeholder="Usuario"></b-input>
                </b-field>
                <b-field label="Contraseña" name="contraseña">
                    <b-input v-model="detalles.contraseña" placeholder="Contraseña" type="password" class="form-control"
                        required></b-input>
                </b-field>
                <b-field label="Confirmar contraseña">
                    <b-input v-model="detalles.confirmarcontraseña" placeholder="Confirmar contraseña" type="password"
                        :locale="locale" inputmode="checkPassword(detalles.contraseña)" required></b-input>
                </b-field>
                <b-field>
                    <input type="submit" value="VALIDAR">
                    <b-button @click="CheckPassworrd()" type="is-success">Guardar</b-button>
                </b-field>
            </form>
        </div>

        <div class="column is-half">
            <ListaUsuario />
        </div>
    </div>
</template>
<script>
import DialogosService from "../services/DialogosService";
import LoginService from "../services/LoginService";
import { DialogProgrammatic as Dialog } from "buefy";
import ListaUsuario from "./ListaUsuario.vue";
export default {

    data: () => ({

        edad: null,

        locale: "es-MX",

        numeric: true,
        checkboxGroup: [],
        detalles: {
            id: null,
            clues: "",
            nombreunidad: "",
            usuario: "",
            contraseña: "",
            confirmarcontraseña: ""
        },
    }),

    methods: {
        
        checkForm2() {
        var valor=this.detalles.contraseña
            var myregex = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/;
            if (myregex.test(valor)) {
                alert(valor + " es valido :-) !");
                return true;
            }
            else {
                alert(valor + " No es valido!");
                return false;
            }
        },
        CheckPassworrd() {
            if (this.detalles.usuario == "") {
                alert("Error: Debe escribir Usuario!");
              //  this.focus();
                return false;
            }
            var res = /^\w+$/;
            if (!res.test(this.detalles.usuario)) {
                alert("Error: Nombre de usuario debe contener únicamente letras, numeros y underscores!");
              //  form.usuario.focus();
                return false;
            }

            if (this.detalles.contraseña != "" && this.detalles.contraseña == this.detalles.confirmarcontraseña) {
                if (!this.CheckPassworrd(this.detalles.contraseña)) {
                    alert("La contraseña no es valida!");
                    this.detalles.contraseña.focus();
                    return false;
                }
            }
            else {
                alert("Error: las contraseñas no coinciden!");
                this.detalles.contraseña.focus();
                return false;
            }
            return true;
        },
        
        checkForm: function (e) {
            this.errors = [];

            if (!this.usuario) {
                this.errors.push("Name required.");
            }
            if (!this.contraseña) {
                this.errors.push('Email required.');
            } else if (!this.validPassword(this.contraseña)) {
                this.errors.push('Valid email required.');
            }

            if (!this.errors.length) {
                return true;
            }

            e.preventDefault();
        },
        validPassword: function (password) {
            var re = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/;
            return re.test(password);
        },
        confirm() {
            Dialog.confirm({
                title: "Guardar registro.",
                message: "Confirmar para guardar datos a la base de datos.",
                cancelText: "Cancelar",
                confirmText: "Aceptar",
                type: "is-success",
                onConfirm: () => this.guardar(),
            });
        },
        async guardar() {
            if (!this.detalles.clues == null) {
                return DialogosService.mostrarNotificacionError("Campos vacios");
            }
            const cargaUtil = {
                nombreunidad: this.detalles.nombreunidad,
                clues: this.detalles.clues,
                usuario: this.detalles.usuario,
                contraseña: this.detalles.contraseña,
                confirmarcontraseña: this.detalles.confirmarcontraseña,
            };
            const respuesta = await LoginService.InsertUsuario(cargaUtil);
            if (respuesta != true) {
                console.log("respuesta save", respuesta);
                DialogosService.mostrarError("registro con errores");
            }
            else {
                DialogosService.mostrarExitoso("Datos guardados correctamente");
                this.detalles = {
                    clues: "",
                    nombreunidad: "",
                    usuario: "",
                    contraseña: "",
                    confirmarcontraseña: "",
                };
            }
        },
    },
    components: { ListaUsuario }
};
</script>