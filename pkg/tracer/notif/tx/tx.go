package tx

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
)

func trace(span trace1.Span, in *npool.TxReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("TxID.%v", index), in.GetTxID()),
		attribute.String(fmt.Sprintf("NotifState.%v", index), in.GetNotifState().String()),
		attribute.String(fmt.Sprintf("TxType.%v", index), in.GetTxType().String()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.TxReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.TxReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
