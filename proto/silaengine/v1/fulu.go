package silaenginev1

func (ebe *ExecutionBundleFulu) GetDecodedSilaRequests(limits ExecutionRequestLimits) (*SilaRequests, error) {
	return decodeExecutionRequestList(ebe.SilaRequests, limits)
}
