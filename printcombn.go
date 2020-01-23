package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n > 0 && n < 10 {
		if n == 1 {
			for i := '0'; i <= '9'; i++ {
				z01.PrintRune(i)
				if i != '8' {
					z01.PrintRune(',')
				}
			}
			z01.PrintRune('\n')
		}
		if n == 2 {
			for i := '0'; i < '9'; i++ {
				for j := i + 1; j <= '9'; j++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					if i != '8' {
						z01.PrintRune(',')
					}

				}
			}
			z01.PrintRune('\n')
		}
		if n == 3 {
			for i := '0'; i < '8'; i++ {
				for j := i + 1; j < '9'; j++ {
					for k := j + 1; k <= '9'; k++ {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						if i != '7' {
							z01.PrintRune(',')
						}
					}

				}
			}
			z01.PrintRune('\n')
		}
		if n == 4 {
			for i := '0'; i < '7'; i++ {
				for j := i + 1; j < '8'; j++ {
					for k := j + 1; k < '9'; k++ {
						for l := k + 1; l <= '9'; l++ {

							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							if i != '6' {
								z01.PrintRune(',')
							}
						}
					}

				}
			}
			z01.PrintRune('\n')
		}
		if n == 5 {
			for i := '0'; i < '6'; i++ {
				for j := i + 1; j < '7'; j++ {
					for k := j + 1; k < '8'; k++ {
						for l := k + 1; l < '9'; l++ {
							for m := l + 1; m <= '9'; m++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(m)
								if i != '5' {
									z01.PrintRune(',')
								}
							}

						}
					}

				}
			}
			z01.PrintRune('\n')
		}
		if n == 6 {
			for i := '0'; i < '5'; i++ {
				for j := i + 1; j < '6'; j++ {
					for k := j + 1; k < '7'; k++ {
						for l := k + 1; l < '8'; l++ {
							for m := l + 1; m < '9'; m++ {
								for n := m + 1; n <= '9'; n++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(m)
									z01.PrintRune(n)
									if i != '4' {
										z01.PrintRune(',')
									}
								}
							}
						}
					}
				}
			}
			z01.PrintRune('\n')
		}
		if n == 7 {
			for i := '0'; i < '4'; i++ {
				for j := i + 1; j < '5'; j++ {
					for k := j + 1; k < '6'; k++ {
						for l := k + 1; l < '7'; l++ {
							for m := l + 1; m < '8'; m++ {
								for n := m + 1; n < '9'; n++ {
									for o := n + 1; n <= '9'; o++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(o)
										if i != '3' {
											z01.PrintRune(',')
										}
									}

								}
							}
						}
					}
				}
			}
			z01.PrintRune('\n')
		}
		if n == 8 {
			for i := '0'; i < '3'; i++ {
				for j := i + 1; j < '4'; j++ {
					for k := j + 1; k < '5'; k++ {
						for l := k + 1; l < '6'; l++ {
							for m := l + 1; m < '7'; m++ {
								for n := m + 1; n < '8'; n++ {
									for o := n + 1; n < '9'; o++ {
										for p := o + 1; p <= '9'; p++ {
											z01.PrintRune(i)
											z01.PrintRune(j)
											z01.PrintRune(k)
											z01.PrintRune(l)
											z01.PrintRune(m)
											z01.PrintRune(n)
											z01.PrintRune(o)
											z01.PrintRune(p)
											if i != '2' {
												z01.PrintRune(',')
											}
										}

									}

								}
							}
						}
					}
				}
			}
			z01.PrintRune('\n')
		}
		if n == 9 {
			for i := '0'; i < '2'; i++ {
				for j := i + 1; j < '3'; j++ {
					for k := j + 1; k < '4'; k++ {
						for l := k + 1; l < '5'; l++ {
							for m := l + 1; m < '6'; m++ {
								for n := m + 1; n < '7'; n++ {
									for o := n + 1; n < '8'; o++ {
										for p := o + 1; p < '9'; p++ {
											for q := p + 1; q <= '9'; q++ {
												z01.PrintRune(i)
												z01.PrintRune(j)
												z01.PrintRune(k)
												z01.PrintRune(l)
												z01.PrintRune(m)
												z01.PrintRune(n)
												z01.PrintRune(o)
												z01.PrintRune(p)
												z01.PrintRune(q)
												if i != '2' {
													z01.PrintRune(',')
												}
											}

										}

									}

								}
							}
						}
					}
				}
			}
			z01.PrintRune('\n')
		}

	}
}

func main() {
	PrintCombN(3)
}
