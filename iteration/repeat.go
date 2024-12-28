package iteration

func Repeat(s string, repeat int) string {
  var repeated string;
  for i := 0; i < repeat; i++ {
    repeated += s;
  }

  return repeated;
}
