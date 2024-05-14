package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := `478819
481713
491341
491341
491341
491341
448164
450766
450766
450766
450766
450766
450766
450766
450766
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
453318
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
458337
469747
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
472832
473820
473820
473820
473820
473820
473820
473820
473820
473820
473820
473820
473820
473820
473820
473959
473959
473959
473959
473959
473959
473959
473959
473959
473959
473959
473959
473959
473959
478008
478008
478008
478008
489080
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491341
491353
491362
491367
491372
453318
453318
453318
453318
489080
489080
491715
478008
492129
492129
492463
421813
421813
421813
421813
421813
421813
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
482490
488603
488604
488605
489039
489039
489039
490493
490493
490493
490493
490493
490493
490493
490493
490493
490493
490493
490493
490493
490493
491715
491715
491715
491715
491715
491715
491715
491715
491715
491715
491715
491715
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492308
492702
492735
492735
492735
492735
492735
492735
492735
492735
492735
492735
492735
492735
492702
492788
492788
492788
492788
492788
492788
492788
492788
492788
492788
492788
492788
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
492810
493067
493067
493067
493074
493074
493074
493074
493074
493074
493074
493074
493074
493074
493074
493074
493089
493089
493089
493284
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
484258
488604
488604
488604
488604
488604
488603
488603
488603
488603
488605
488605
488605
488605
489039
436252
466640
474796
436252
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
480696
436252
485972
489298
490131
490131
490131
490131
490131
490131
491425
491973
491973
491973
491973
491973
491973
488599
488599
488599
488599
488599
488599
488599
488599
488599
488599
488599
488599
494039
490131
490131
490131
490131
490131
490131
494058
494058
494058
494066
494066
494066
494202
494202
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494224
494568
494568
494607
494611
495240
495240
495240
495240
495325
495626
495626
495626
495626
379077
379077
396362
396362
396362
397336
397336
397336
397336
397423
397428
397428
397428
397428
396362
397423
397423
404120
404120
404120
404120
404120
404120
412887
418036
418036
418036
418036
418036
418036
418036
418036
423323
423323
426230
426230
427216
431501
431501
431770
431770
450000
450008
423323
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465340
465395
465395
470425
470425
470425
470425
470425
470425
470425
470425
470425
470425
471631
471631
473617
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475797
475807
475810
478034
478034
478034
478034
479090
479090
479090
479090
479090
479090
479090
479090
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
479208
473617
492113
492113
492137
492137
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
465088
492573
492573
492573
492573
470425
470425
470425
470425
470425
492679
492679
492679
492679
492679
492679
492679
492679
492694
465395
465395
465395
465395
465395
493610
495148
495469
495469
495469
495469
496215
496782
496782
496782
496782
496842
496842
496842
496842
496842
496842
496842
496842
496842
496842
496842
496842
496848
496848
496848
496848
496848
496848
496848
496848
496848
496848
496848
496848
497170
497170
497170
497170
497170
497170
497170
497170
497240
497240
497240
493610
493610
493610
493610
493610
431501
497883
497883
`

	// 使用strings.Split将字符串按换行符分割成切片
	lines := strings.Split(numbers, "\n")

	// 使用strings.Builder来构建结果字符串
	var sb strings.Builder

	// 遍历切片，并在每个数字后面添加逗号，除了最后一个
	for i, line := range lines {
		sb.WriteString(line)
		if i < len(lines)-1 {
			sb.WriteString(", ")
		}
	}

	// 获取构建的结果字符串
	result := sb.String()

	// 打印结果
	fmt.Println("test分支", result)
	fmt.Println("test1", result)
}
