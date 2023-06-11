import HttpService from "./HttpService";
const VehiculosService = {
    async insertEmbarazada(embarazada) {
        console.log("frontend",embarazada);
        return await HttpService.post("/agregar_embarazada", embarazada);
    },
    async GetEmbarazada() {
        return await HttpService.get("/Lista_embarazada");
    },
    async obtenerMunicipios(){
        return await HttpService.get("/agregar_embarazada");
       
    },
    async obteneLocalidadesById(idMunicipio) {
        return await HttpService.get(`/agregar_embarazada/${idMunicipio}`);
    },
    async agregarPuerpera(vehiculo) {
        return await HttpService.post("/agregar_puerpera", vehiculo);
    },

};
export default VehiculosService;