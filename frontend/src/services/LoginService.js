import HttpService from "./HttpService";
const LoginService = {
    async InsertUsuario(login) {
        return await HttpService.post("/Tarjetas", login);
    },
    
    async obtenerUsuarios() {
        return await HttpService.get("/ListaUsuario");
    },

};
export default LoginService;