package sigdev

import (
	"strconv"
)

func processByteSequences(a *SignatureInterface) ([]byteSequence, error) {
	subSequences := make([]byteSequence, len(a.Sequences))
	for idx, sig := range a.Sequences {
		var signatureSequence byteSequence
		signatureSequence.Sequence = sig.Sequence
		if sig.Relativity == VAR || sig.Relativity == VAROFF || sig.Relativity == "" {
			subSequences[idx] = signatureSequence
			continue

		}
		signatureSequence.MinOffset = strconv.Itoa(sig.Offset)
		signatureSequence.MaxOffset = strconv.Itoa(sig.MaxOffset)
		signatureSequence.Reference = sig.Relativity
		subSequences[idx] = signatureSequence
	}
	return subSequences, nil
}
