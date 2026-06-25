package silaenginev1

func (ebe *ExecutionBundleGloas) GetDecodedSilaRequests(limits ExecutionRequestLimits) (*SilaRequests, error) {
	return decodeExecutionRequestList(ebe.SilaRequests, limits)
}
