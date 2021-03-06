// Code generated by goa v3.7.2, DO NOT EDIT.
//
// inventory-system HTTP client CLI support package
//
// Command:
// $ goa gen inventory-system/api/svc/design

package cli

import (
	"flag"
	"fmt"
	inventoryc "inventory-system/gen/http/inventory/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `inventory (create|update|find|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` inventory create --body '{
      "productDescription": "Voluptatum maxime debitis soluta.",
      "productMinStock": 1041205383,
      "productName": "Occaecati consequuntur amet inventore."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		inventoryFlags = flag.NewFlagSet("inventory", flag.ContinueOnError)

		inventoryCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		inventoryCreateBodyFlag = inventoryCreateFlags.String("body", "REQUIRED", "")

		inventoryUpdateFlags         = flag.NewFlagSet("update", flag.ExitOnError)
		inventoryUpdateBodyFlag      = inventoryUpdateFlags.String("body", "REQUIRED", "")
		inventoryUpdateProductIDFlag = inventoryUpdateFlags.String("product-id", "REQUIRED", "")

		inventoryFindFlags         = flag.NewFlagSet("find", flag.ExitOnError)
		inventoryFindProductIDFlag = inventoryFindFlags.String("product-id", "REQUIRED", "")

		inventoryDeleteFlags         = flag.NewFlagSet("delete", flag.ExitOnError)
		inventoryDeleteProductIDFlag = inventoryDeleteFlags.String("product-id", "REQUIRED", "")
	)
	inventoryFlags.Usage = inventoryUsage
	inventoryCreateFlags.Usage = inventoryCreateUsage
	inventoryUpdateFlags.Usage = inventoryUpdateUsage
	inventoryFindFlags.Usage = inventoryFindUsage
	inventoryDeleteFlags.Usage = inventoryDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "inventory":
			svcf = inventoryFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "inventory":
			switch epn {
			case "create":
				epf = inventoryCreateFlags

			case "update":
				epf = inventoryUpdateFlags

			case "find":
				epf = inventoryFindFlags

			case "delete":
				epf = inventoryDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "inventory":
			c := inventoryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = inventoryc.BuildCreatePayload(*inventoryCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = inventoryc.BuildUpdatePayload(*inventoryUpdateBodyFlag, *inventoryUpdateProductIDFlag)
			case "find":
				endpoint = c.Find()
				data, err = inventoryc.BuildFindPayload(*inventoryFindProductIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = inventoryc.BuildDeletePayload(*inventoryDeleteProductIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// inventoryUsage displays the usage of the inventory command and its
// subcommands.
func inventoryUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers.
Usage:
    %[1]s [globalflags] inventory COMMAND [flags]

COMMAND:
    create: Create implements create.
    update: Update implements update.
    find: Find implements find.
    delete: Delete implements delete.

Additional help:
    %[1]s inventory COMMAND --help
`, os.Args[0])
}
func inventoryCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s inventory create --body '{
      "productDescription": "Voluptatum maxime debitis soluta.",
      "productMinStock": 1041205383,
      "productName": "Occaecati consequuntur amet inventore."
   }'
`, os.Args[0])
}

func inventoryUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory update -body JSON -product-id INT

Update implements update.
    -body JSON: 
    -product-id INT: 

Example:
    %[1]s inventory update --body '{
      "productDescription": "Itaque ut sunt.",
      "productMinStock": 136779013,
      "productName": "Sint rerum ab."
   }' --product-id 8948898044999578082
`, os.Args[0])
}

func inventoryFindUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory find -product-id INT

Find implements find.
    -product-id INT: 

Example:
    %[1]s inventory find --product-id 3544923032055856752
`, os.Args[0])
}

func inventoryDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory delete -product-id INT

Delete implements delete.
    -product-id INT: 

Example:
    %[1]s inventory delete --product-id 77523483320048548
`, os.Args[0])
}
