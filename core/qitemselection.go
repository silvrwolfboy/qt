package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QItemSelection struct {
	QList
}

type QItemSelection_ITF interface {
	QList_ITF
	QItemSelection_PTR() *QItemSelection
}

func PointerFromQItemSelection(ptr QItemSelection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemSelection_PTR().Pointer()
	}
	return nil
}

func NewQItemSelectionFromPointer(ptr unsafe.Pointer) *QItemSelection {
	var n = new(QItemSelection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemSelection) QItemSelection_PTR() *QItemSelection {
	return ptr
}

func NewQItemSelection() *QItemSelection {
	defer qt.Recovering("QItemSelection::QItemSelection")

	return NewQItemSelectionFromPointer(C.QItemSelection_NewQItemSelection())
}

func NewQItemSelection2(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) *QItemSelection {
	defer qt.Recovering("QItemSelection::QItemSelection")

	return NewQItemSelectionFromPointer(C.QItemSelection_NewQItemSelection2(PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight)))
}

func (ptr *QItemSelection) Contains(index QModelIndex_ITF) bool {
	defer qt.Recovering("QItemSelection::contains")

	if ptr.Pointer() != nil {
		return C.QItemSelection_Contains(ptr.Pointer(), PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QItemSelection) Merge(other QItemSelection_ITF, command QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QItemSelection::merge")

	if ptr.Pointer() != nil {
		C.QItemSelection_Merge(ptr.Pointer(), PointerFromQItemSelection(other), C.int(command))
	}
}

func (ptr *QItemSelection) Select(topLeft QModelIndex_ITF, bottomRight QModelIndex_ITF) {
	defer qt.Recovering("QItemSelection::select")

	if ptr.Pointer() != nil {
		C.QItemSelection_Select(ptr.Pointer(), PointerFromQModelIndex(topLeft), PointerFromQModelIndex(bottomRight))
	}
}

func QItemSelection_Split(ran QItemSelectionRange_ITF, other QItemSelectionRange_ITF, result QItemSelection_ITF) {
	defer qt.Recovering("QItemSelection::split")

	C.QItemSelection_QItemSelection_Split(PointerFromQItemSelectionRange(ran), PointerFromQItemSelectionRange(other), PointerFromQItemSelection(result))
}