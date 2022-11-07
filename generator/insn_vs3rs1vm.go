package generator

import (
	"fmt"
	"strconv"
	"strings"
)

func (i *insn) genCodeVs3Rs1mVm() string {
	getEEW := func(name string) SEW {
		eew, _ := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(i.Name, "vse"), ".v"))
		return SEW(eew)
	}

	builder := strings.Builder{}
	builder.WriteString(i.gTestDataAddr())
	builder.WriteString(i.gWriteRandomData(LMUL(1)))
	builder.WriteString(i.gLoadDataIntoRegisterGroup(0, LMUL(1), SEW(8)))

	for _, c := range i.combinations([]SEW{getEEW(i.Name)}) {
		builder.WriteString(c.comment())

		vs3 := int(c.LMUL1)
		builder.WriteString(i.gWriteTestData(c.LMUL1, c.SEW, 0))
		builder.WriteString(i.gLoadDataIntoRegisterGroup(vs3, c.LMUL1, c.SEW))
		builder.WriteString(i.gWriteRandomData(c.LMUL1))

		builder.WriteString("# -------------- TEST BEGIN --------------\n")
		builder.WriteString(i.gVsetvli(c.Vl, c.SEW, c.LMUL))
		builder.WriteString(fmt.Sprintf("%s v%d, (a0)%s\n", i.Name, vs3, v0t(c.Mask)))
		builder.WriteString("# -------------- TEST END   --------------\n")

		builder.WriteString(i.gLoadDataIntoRegisterGroup(vs3, c.LMUL1, c.SEW))
		builder.WriteString(i.gMagicInsn(vs3))
	}
	return builder.String()
}
