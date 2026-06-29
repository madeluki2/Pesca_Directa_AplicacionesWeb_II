// Tests del almacén en memoria con la librería estándar (testing), sin testify.
// Es el estilo más básico de Go: t.Fatalf para abortar, t.Errorf para seguir.
package storage

import (
	"testing"

	"Pesca_Directa_AplicacionesWeb_II/internal/models"
)

func TestMemoriaPedidos_CrearYBuscarCliente(t *testing.T) {
	m := NewMemoria()

	creado := m.CrearCliente(models.Cliente{
		NombreNegocio: "Sushi Koi",
		TipoCliente:   "restaurante",
		Direccion:     "Av. Flavio Reyes",
		Telefono:      "0991234567",
		Estado:        "activo",
	})
	if creado.ID == 0 {
		t.Fatalf("esperaba un ID asignado, obtuve 0")
	}

	encontrado, ok := m.BuscarClientePorID(creado.ID)
	if !ok {
		t.Fatalf("no se encontró el cliente recién creado (id=%d)", creado.ID)
	}
	if encontrado.NombreNegocio != "Sushi Koi" {
		t.Errorf("nombre = %q; esperaba %q", encontrado.NombreNegocio, "Sushi Koi")
	}
}

func TestMemoriaPedidos_BuscarClienteInexistente(t *testing.T) {
	m := NewMemoria()

	// El patrón comma-ok: ok debe ser false para un id que no existe.
	if _, ok := m.BuscarClientePorID(999); ok {
		t.Errorf("esperaba ok=false para un id inexistente")
	}
}

func TestMemoriaPedidos_ActualizarYEliminarCliente(t *testing.T) {
	m := NewMemoria()
	creado := m.CrearCliente(models.Cliente{
		NombreNegocio: "Distribuidora El Puerto",
		TipoCliente:   "mayorista",
		Direccion:     "Calle 10 de Agosto",
		Telefono:      "0997654321",
		Estado:        "activo",
	})

	if _, ok := m.ActualizarCliente(creado.ID, models.Cliente{NombreNegocio: "Distribuidora Nueva"}); !ok {
		t.Fatalf("no se pudo actualizar el cliente id=%d", creado.ID)
	}

	if !m.EliminarCliente(creado.ID) {
		t.Errorf("esperaba poder eliminar el cliente id=%d", creado.ID)
	}
	if _, ok := m.BuscarClientePorID(creado.ID); ok {
		t.Errorf("el cliente id=%d debería haber sido eliminado", creado.ID)
	}
}
