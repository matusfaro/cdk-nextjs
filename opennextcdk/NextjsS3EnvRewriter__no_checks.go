//go:build no_runtime_type_checking

package opennextcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateNextjsS3EnvRewriter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNextjsS3EnvRewriterParameters(scope constructs.Construct, id *string, props *NextjsS3EnvRewriterProps) error {
	return nil
}
