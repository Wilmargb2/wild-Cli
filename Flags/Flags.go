package flags

import "flag"

type Flags struct {
	Make     bool
	CompName string
	HookName string
	ServiceName string
}

func InitFlags() *Flags {

	//Make contempla las siguientes flags
	make := flag.Bool("make", false, "Crear comps, hooks, services --make")

	// Flags para crear recursos con --make
	CompName := flag.String(
		"comp",
		"",
		"Nombre del componente (úsalo con --make). Ej: --comp=Card",
	)

	HookName := flag.String(
		"hook",
		"",
		"Nombre del hook sin 'use o hook' (úsalo con --make).",
	)

	ServiceName := flag.String(
		"service",
		"",
		"Nombre del service sin 'use o service' (úsalo con --make).",
	)

	flag.Parse() //--> Crea las flags, para su uso

	return &Flags{
		Make:     *make,
		CompName: *CompName,
		HookName: *HookName,
		ServiceName: *ServiceName,
	}
}
