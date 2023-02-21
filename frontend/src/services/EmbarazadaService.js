import HttpService from "./HttpService";
const VehiculosService = {
    async insertEmbarazada(embarazada) {
        console.log("frontend",embarazada);
        return await HttpService.post("/agregar_embarazada", embarazada);
    },
    async obtenerEmbarazadas() {
        return await HttpService.get("/ListaEmbarazada");
    },
    async obtenerMunicipios(){
        return await HttpService.get("/agregar_embarazada");
       
    },
    async obteneLocalidadesById(idMunicipio) {
        return await HttpService.get(`/agregar_embarazada/${idMunicipio}`);
    },

};
export default VehiculosService;