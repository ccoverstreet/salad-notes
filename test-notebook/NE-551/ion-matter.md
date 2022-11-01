# Ion-Matter Interactions
2022-09-15

$2\AA$

salskjdalkd

asdasds

asdas


asdas
asdasda

sd


![asd](./Figure_4.svg)

asd

$$
	S_c = 4 \pi r_0^2 m_e c^2 (\frac{z^2}{\beta^2})(\frac{N_A\rho Z}{M_m})
	\{ln(\frac{2m_e c^2 \gamma^2 \beta^2}{I}) - \beta^2\}
$$ 


\begin{gather*}
	\alpha \rightarrow T/A=200\ \text{MeV},\ T=800\ \text{MeV} \\ 
	r_0 = 2.818\times 10^{-13}\ \text{cm} \\
	m_e c^2 = 0.511\ \text{MeV} \\ 
	z^2 = 4 \\\\
	E_\alpha = T_\alpha + m_\alpha c^2 = \gamma(m_\alpha c^2) \\
	T_\alpha = (\gamma - 1) m_\alpha c^2 \\ \\
	\text{Calculate mass of alpha:} \\
	m_\alpha = (4*931.494\ \text{MeV}) + \delta - 2 m_e = 3727.379 \\\\
	\gamma = 1.215 \\
	\gamma^2 = \frac{1}{1-\beta^2} \\
	\beta^2 = 1-\frac{1}{\gamma^2} = 0.323 \\
	\beta^2 \gamma^2 = \gamma^2 - 1 \\\\
	n_e = \frac{N_A \rho Z}{M_m} \\
	I_{Al} = 166\ \text{eV} \\\\
	\text{Plug into massive equation. Check units} \\
	S_c = 37.908\ \text{(MeV/cm from star database)} \\
	S_c = 37.879\ \text{MeV/cm from Python calc}
\end{gather*}

## Calculating densities

- Use NIST values
	- Air:
		- Z = 6, 7, 8, 18
		- $\rho=1.20479\times 10^-2\ g/cm^3$ 
		- Mean excitation energy = 86.7 eV

\begin{gather*}
	\rho_A = \frac{N_A}{M_m} \rho \\
	\rho_A = \frac{N_A}{mw}\rho n_i \\
	mw = \sum_{i=1}^N = n_i A_i \\
	f_i = \frac{n_i A_i}{\sum_{i=1}^N n_i A_i} = \frac{n_i A_i}{mw} \\ 
	\rho_{A_i} = N_A / (mw) \rho \frac{f_i mw_i}{A_i} \\
	\rho_{A_i} = N_A \rho \frac{f_i}{A_i} \\
	\text{or} \\
	\rho_{A_i} = \frac{N_a \rho f_i}{A_i} \\ 
	\text{Electron density} \\
	n_e = \sum_{i=1}^{N} \frac{N_a \rho f_i}{A_i}Z_i\\
\end{gather*}
	- Don't need molar weight of molecule


- Electron density of water
\begin{gather*}
	A_O = 16 \\
	A_H = 1 \\
	Z_O = 8 \\
	Z_H = 1 \\
	f_H = 2/18 \\ 
	f_O = 16/18 \\
	n_e = N_A \rho [\frac{f_H}{A_H} Z_H + \frac{f_O}{A_O} Z_O] \\ 
	n_e = 3.35\times 10^{23}\ [e^-/cm^3]
\end{gather*}


# END


\begin{gather*}
	2 + 2 \\
	3 \\
\end{gather*}
